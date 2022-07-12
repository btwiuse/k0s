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
    title="BarGauges",
    uid=__file__.split('/')[-1:][0].replace('.py', ''),
    rows=[
      Row(panels=[
        BarGauge(
          title=f"{instance}",
          dataSource='Prometheus',
          targets=[
            Target(
                expr=f'(((count(count(node_cpu_seconds_total{{instance=~"{instance}",job=~"all"}}) by (cpu))) - avg(sum by (mode)(irate(node_cpu_seconds_total{{mode="idle",instance=~"{instance}",job=~"all"}}[5m])))) * 100) / count(count(node_cpu_seconds_total{{instance=~"{instance}",job=~"all"}}) by (cpu))',
                refId='A',
                legendFormat='CPU',
                instant=True,
            ),
            Target(
                expr=f'100 - ((node_memory_MemAvailable_bytes{{instance=~"{instance}",job=~"all"}} * 100) / node_memory_MemTotal_bytes{{instance=~"{instance}",job=~"all"}})',
                legendFormat='Memory',
                instant=True,
            ),
            Target(
                expr=f'100 - ((node_filesystem_avail_bytes{{instance=~"{instance}",job=~"all",mountpoint="/",fstype!="rootfs"}} * 100) / node_filesystem_size_bytes{{instance=~"{instance}",job=~"all",mountpoint="/",fstype!="rootfs"}})',
                legendFormat='Disk',
                instant=True,
            ),
          ],
        ),
      ]) for instance in nodes
    ],
).auto_panel_ids()
