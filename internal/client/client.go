// Package client is a client library for interacting with a Transiter server.
//
// It wraps the raw gRPC API with a more idiomatic interface.
package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jamespfennell/transiter/internal/client/table"
	"github.com/jamespfennell/transiter/internal/gen/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn         *grpc.ClientConn
	publicClient api.PublicClient
	adminClient  api.AdminClient
}

func New(addr string) (*Client, error) {
	var err error
	// TODO: credentials?
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return &Client{
		conn:         conn,
		publicClient: api.NewPublicClient(conn),
		adminClient:  api.NewAdminClient(conn),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) DeleteSystem(ctx context.Context, systemID string) error {
	req := api.DeleteSystemRequest{SystemId: systemID}
	_, err := c.adminClient.DeleteSystem(ctx, &req)
	return err
}

type ConfigPathType int

const (
	TransiterRepo ConfigPathType = iota
	File
	URL
)

type InstallSystemArgs struct {
	SystemID       string
	ConfigPath     string
	ConfigPathType ConfigPathType
	AllowUpdate    bool
	IsTemplate     bool
	TemplateArgs   map[string]string
}

func (c *Client) InstallSystem(ctx context.Context, args InstallSystemArgs) error {
	yamlContent, err := getYamlContent(args)
	if err != nil {
		return err
	}
	req := api.InstallOrUpdateSystemRequest{
		SystemId:    args.SystemID,
		InstallOnly: !args.AllowUpdate,
		Config: &api.InstallOrUpdateSystemRequest_YamlConfig{
			YamlConfig: &api.YamlConfig{
				Content:      yamlContent,
				IsTemplate:   args.IsTemplate,
				TemplateArgs: args.TemplateArgs,
			},
		},
	}
	_, err = c.adminClient.InstallOrUpdateSystem(ctx, &req)
	if err != nil {
		return err
	}
	fmt.Println("System install started; waiting for system to become active.")
	for {
		system, err := c.publicClient.GetSystem(ctx, &api.GetSystemRequest{SystemId: args.SystemID})
		if err != nil {
			return fmt.Errorf("failed to poll system status: %w", err)
		}
		switch system.Status {
		case api.System_ACTIVE:
			fmt.Println("System installed.")
			return nil
		case api.System_INSTALL_FAILED, api.System_UPDATE_FAILED:
			return fmt.Errorf("failed to install/update system")
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func getYamlContent(args InstallSystemArgs) (string, error) {
	switch args.ConfigPathType {
	case File:
		fmt.Printf("Reading system config from file %s.\n", args.ConfigPath)
		yaml, err := os.ReadFile(args.ConfigPath)
		if err != nil {
			return "", err
		}
		return string(yaml), nil
	case URL:
		fmt.Printf("Getting system config from URL %s.\n", args.ConfigPath)
		yaml, err := getURL(args.ConfigPath)
		if err != nil {
			return "", fmt.Errorf("failed to read transit system config from URL: %w", err)
		}
		return yaml, nil
	default:
		url := fmt.Sprintf("https://raw.githubusercontent.com/jamespfennell/transiter/master/systems/%s.yaml", args.ConfigPath)
		fmt.Printf("Getting system config from Transiter repository at %s.\n", url)
		yaml, err := getURL(url)
		if err != nil {
			return "", fmt.Errorf("failed to read transit system config from the Transiter repository: %w", err)
		}
		return yaml, nil
	}
}

func getURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (c *Client) UpdateFeed(ctx context.Context, systemID, feedID string, force bool) error {
	resp, err := c.adminClient.UpdateFeed(ctx, &api.UpdateFeedRequest{
		SystemId: systemID,
		FeedId:   feedID,
		Force:    force,
	})
	if err != nil {
		return err
	}
	switch resp.GetFeedUpdate().GetStatus() {
	case api.FeedUpdate_SKIPPED:
		fmt.Println("Feed update skipped because downloaded data was identical to the last update (use --force to disable this check)")
		return nil
	case api.FeedUpdate_UPDATED:
		fmt.Printf("Feed updated in %s\n", time.Duration(resp.GetFeedUpdate().GetTotalLatencyMs())*time.Millisecond)
		return nil
	default:
		return fmt.Errorf("feed update failed with reason %s: %s", resp.GetFeedUpdate().GetStatus(), resp.GetFeedUpdate().GetErrorMessage())
	}
}

func (c *Client) ListSystems(ctx context.Context) error {
	var req api.ListSystemsRequest
	rep, err := c.publicClient.ListSystems(ctx, &req)
	if err != nil {
		return err
	}
	if len(rep.Systems) == 0 {
		fmt.Println("No transit systems installed.")
		return nil
	}
	t := table.New()
	t.AddRow("ID", "Name", "Status")
	t.AddSeperator()
	for _, system := range rep.Systems {
		t.AddRow(system.Id, system.Name, system.Status.String())
	}
	fmt.Printf("%s", t.Render())
	return nil
}

func (c *Client) SchedulerStatus(ctx context.Context) error {
	var req api.GetSchedulerStatusRequest
	reply, err := c.adminClient.GetSchedulerStatus(ctx, &req)
	if err != nil {
		return err
	}
	t := table.New()
	t.AddRow("System ID", "Feed ID", "Scheduling policy", "Last", "Last", "Currently")
	t.AddRow("", "", "", "finished", "successful", "running")
	t.AddRow("", "", "", "update", "update", "")
	t.AddSeperator()
	var lastSystemID string
	for _, feed := range reply.Feeds {
		if lastSystemID != "" && lastSystemID != feed.SystemId {
			t.AddSeperator()
		}
		var systemIDToPrint string
		if lastSystemID == "" || lastSystemID != feed.SystemId {
			systemIDToPrint = feed.SystemId
		}
		lastSystemID = feed.SystemId
		var schedulingPolicy string
		switch feed.FeedConfig.GetSchedulingPolicy() {
		case api.FeedConfig_DAILY:
			schedulingPolicy = fmt.Sprintf(
				"daily @ %s (%s)",
				feed.FeedConfig.GetDailyUpdateTime(),
				feed.FeedConfig.GetDailyUpdateTimezone())
		case api.FeedConfig_PERIODIC:
			schedulingPolicy = fmt.Sprintf(
				"periodic %s",
				(time.Millisecond * time.Duration(feed.FeedConfig.GetPeriodicUpdatePeriodMs())).String())
		default:
			schedulingPolicy = "unknown"
		}
		t.AddRow(
			systemIDToPrint,
			feed.FeedConfig.GetId(),
			schedulingPolicy,
			convertTime(feed.LastFinishedUpdate),
			convertTime(feed.LastSuccessfulUpdate),
			fmt.Sprintf("%t", feed.CurrentlyRunning))
	}
	fmt.Printf("%s", t.Render())
	return nil
}

func (c *Client) ResetScheduler(ctx context.Context) error {
	_, err := c.adminClient.ResetScheduler(ctx, &api.ResetSchedulerRequest{})
	return err
}

func convertTime(t int64) string {
	if t == 0 {
		return "(none)"
	}
	return fmt.Sprintf("%s ago", time.Now().Round(time.Second).Sub(time.Unix(t, 0)))
}

func (c *Client) GetLogLevel(ctx context.Context) error {
	r, err := c.adminClient.GetLogLevel(ctx, &api.GetLogLevelRequest{})
	if err != nil {
		return err
	}
	fmt.Println(r.GetLogLevel())
	return nil
}

func (c *Client) SetLogLevel(ctx context.Context, logLevel string) error {
	ll, ok := api.LogLevel_value[logLevel]
	if !ok {
		return fmt.Errorf("invalid log level %q; allowed values are debug, info, warn, error", logLevel)
	}
	_, err := c.adminClient.SetLogLevel(ctx, &api.SetLogLevelRequest{
		LogLevel: api.LogLevel(ll),
	})
	return err
}

func (c *Client) Version(ctx context.Context) (string, error) {
	r, err := c.publicClient.Entrypoint(ctx, &api.EntrypointRequest{})
	return r.GetTransiter().GetVersion(), err
}
