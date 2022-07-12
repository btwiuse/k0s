import yaml
from grafanalib.core import (
    Alert, AlertCondition, Dashboard, Graph,
    GreaterThan, OP_AND, OPS_FORMAT, Row, Column, RTYPE_SUM, SECONDS_FORMAT,
    SHORT_FORMAT, single_y_axis, Target, TimeRange, YAxes, YAxis
)

nodes=[]

with open("nodelist.yaml", 'r') as stream:
    nodes = yaml.safe_load(stream)

dashboard = Dashboard(
    title="Agents",
    uid=__file__.split('/')[-1:][0].replace('.py', ''),
    rows=[
      Row(panels=[
        Graph(
          title="{}".format(i),
          dataSource='Prometheus',
          targets=[
            Target(
                expr='process_resident_memory_bytes{{job="all",instance="{}"}}'.format(i),
                refId='A',
            ),
          ],
        )
      ]) for i in nodes
    ],
).auto_panel_ids()
