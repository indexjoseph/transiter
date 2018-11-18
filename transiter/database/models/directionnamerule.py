from sqlalchemy import Index, Column, Integer, String, Boolean, ForeignKey
from sqlalchemy.orm import relationship

from .base import Base


class DirectionNameRule(Base):
    __tablename__ = 'direction_name_rule'

    pk = Column(Integer, primary_key=True)
    stop_pk = Column(Integer, ForeignKey('stop.pk'))

    priority = Column(Integer)
    direction_id = Column(Boolean)
    track = Column(String)
    stop_id_alias = Column(String)
    name = Column(String)

    stop = relationship(
        'Stop',
        back_populates='direction_name_rules')


Index('direction_name_rule_stop_pk_priority_idx',
      DirectionNameRule.stop_pk,
      DirectionNameRule.priority)

