# MysqlGate

Mostly inspired by PostgreSQL, MysqlGate is another proxy for mysql but aimed at extendable.


## Designed Extend Point

- User defined functions/aggregates
- Background worker processes
  - Example usage: timescaledb like plugin
- Query planner. 
  - Example usage: sharding/ timescaledb like plugin
- FDW.