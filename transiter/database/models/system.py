from sqlalchemy import Column, Integer, String
from sqlalchemy.orm import relationship

from .base import Base


class System(Base):
    __tablename__ = 'system'

    pk = Column(Integer, primary_key=True)
    id = Column(String, unique=True, index=True)

    name = Column(String, nullable=True)
    package = Column(String, nullable=False)
    directory_hash = Column(String, nullable=True)

    routes = relationship(
        'Route',
        back_populates='system',
        cascade='all, delete-orphan')
    stations = relationship(
        'Station',
        back_populates='system',
        cascade='all, delete-orphan')
    feeds = relationship(
        'Feed',
        back_populates='system',
        cascade='all, delete-orphan')

    _short_repr_list = ['id']
