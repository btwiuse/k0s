import yaml
from grafanalib.core import (
    Alert, AlertCondition, Dashboard, Graph, BarGauge, GaugePanel,
    GreaterThan, OP_AND, OPS_FORMAT, Row, Column, RTYPE_SUM, SECONDS_FORMAT,
    SHORT_FORMAT, single_y_axis, Target, TimeRange, YAxes, YAxis
)

nodes=[]

with open("nodelist.yaml", 'r') as stream:
    nodes = yaml.safe_load(stream)

dashboard = Dashboard(
    title="BarGauges Memory",
    uid=__file__.split('/')[-1:][0].replace('.py', ''),
    rows=[
      Row(panels=[
        BarGauge(
          title=f"Memory",
          height=1200,
          span=12,
          dataSource='Prometheus',
          targets=[
            Target(
                expr=f'100 - ((node_memory_MemAvailable_bytes{{instance=~"{instance}",job=~"all"}} * 100) / node_memory_MemTotal_bytes{{instance=~"{instance}",job=~"all"}})',
                legendFormat=f'{instance}',
                instant=True,
            ) for instance in nodes
          ],
        ),
      ]),
    ],
).auto_panel_ids()
