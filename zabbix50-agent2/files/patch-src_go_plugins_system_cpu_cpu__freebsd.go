--- src/go/plugins/system/cpu/cpu_freebsd.go.orig	2020-05-29 14:32:25.625708000 +0200
+++ src/go/plugins/system/cpu/cpu_freebsd.go	2020-05-29 14:31:44.298984000 +0200
@@ -0,0 +1,136 @@
+/*
+** Zabbix
+** Copyright (C) 2001-2020 Zabbix SIA
+**
+** This program is free software; you can redistribute it and/or modify
+** it under the terms of the GNU General Public License as published by
+** the Free Software Foundation; either version 2 of the License, or
+** (at your option) any later version.
+**
+** This program is distributed in the hope that it will be useful,
+** but WITHOUT ANY WARRANTY; without even the implied warranty of
+** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
+** GNU General Public License for more details.
+**
+** You should have received a copy of the GNU General Public License
+** along with this program; if not, write to the Free Software
+** Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
+**/
+
+package cpu
+
+import (
+	"math"
+
+	"zabbix.com/pkg/plugin"
+
+	"github.com/shirou/gopsutil/cpu"
+)
+
+// Plugin -
+type Plugin struct {
+	plugin.Base
+	cpus []*cpuUnit
+}
+
+func (p *Plugin) getCpuLoad(params []string) (result interface{}, err error) {
+	return nil, plugin.UnsupportedMetricError
+}
+
+func (p *Plugin) Collect() (err error) {
+        cpustat, _ := cpu.Times(true)
+
+	// *sigh* cpu[0] is not cpu 0, it's used for the overall stats
+	// *brilliant* design, thank you
+	cpu := p.cpus[0]
+	cpu.status = cpuStatusOnline
+
+	var user, system, idle, nice, iowait, irq, softirq, steal, guest, guestNice float64
+	for _, stat := range cpustat {
+		user	+= stat.User
+		system	+= stat.System
+		idle	+= stat.Idle
+		nice	+= stat.Nice
+		iowait	+= stat.Iowait
+		irq	+= stat.Irq
+		softirq	+= stat.Softirq
+		steal	+= stat.Steal
+		guest	+= stat.Guest
+		guestNice += stat.GuestNice
+	}
+
+	slot := &cpu.history[cpu.tail]
+	slot.counters[counterUser] =	uint64(math.Floor(user))
+	slot.counters[counterNice] =	uint64(math.Floor(nice))
+	slot.counters[counterSystem] =	uint64(math.Floor(system))
+	slot.counters[counterIdle] =	uint64(math.Floor(idle))
+	slot.counters[counterIowait] =	uint64(math.Floor(iowait))
+	slot.counters[counterIrq] =	uint64(math.Floor(irq))
+	slot.counters[counterSoftirq] =	uint64(math.Floor(softirq))
+	slot.counters[counterSteal] =	uint64(math.Floor(steal))
+	slot.counters[counterGcpu] =	uint64(math.Floor(guest))
+	slot.counters[counterGnice] =	uint64(math.Floor(guestNice))
+
+	if cpu.tail = cpu.tail.inc(); cpu.tail == cpu.head {
+		cpu.head = cpu.head.inc()
+	}
+
+	for index, stat := range cpustat {
+		cpu := p.cpus[index+1]
+		cpu.status = cpuStatusOnline
+
+		slot := &cpu.history[cpu.tail]
+		slot.counters[counterUser] =	uint64(math.Floor(stat.User))
+		slot.counters[counterNice] =	uint64(math.Floor(stat.Nice))
+		slot.counters[counterSystem] =	uint64(math.Floor(stat.System))
+		slot.counters[counterIdle] =	uint64(math.Floor(stat.Idle))
+		slot.counters[counterIowait] =	uint64(math.Floor(stat.Iowait))
+		slot.counters[counterIrq] =	uint64(math.Floor(stat.Irq))
+		slot.counters[counterSoftirq] =	uint64(math.Floor(stat.Softirq))
+		slot.counters[counterSteal] =	uint64(math.Floor(stat.Steal))
+		slot.counters[counterGcpu] =	uint64(math.Floor(stat.Guest))
+		slot.counters[counterGnice] =	uint64(math.Floor(stat.GuestNice))
+
+		if cpu.tail = cpu.tail.inc(); cpu.tail == cpu.head {
+			cpu.head = cpu.head.inc()
+		}
+	}
+	return nil
+}
+
+func numCPU() int {
+	cpuNum, _ := cpu.Counts(true)
+	return cpuNum
+}
+
+func (p *Plugin) Start() {
+	p.cpus = p.newCpus(numCPU())
+}
+
+func (p *Plugin) Stop() {
+	p.cpus = nil
+}
+
+func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {
+	if p.cpus == nil || p.cpus[0].head == p.cpus[0].tail {
+		// no data gathered yet
+		return
+	}
+	switch key {
+	case "system.cpu.discovery":
+		return p.getCpuDiscovery(params)
+	case "system.cpu.num":
+		return p.getCpuNum(params)
+	case "system.cpu.util":
+		return p.getCpuUtil(params)
+	default:
+		return nil, plugin.UnsupportedMetricError
+	}
+}
+
+func init() {
+	plugin.RegisterMetrics(&impl, pluginName,
+		"system.cpu.discovery", "List of detected CPUs/CPU cores, used for low-level discovery.",
+		"system.cpu.num", "Number of CPUs.",
+		"system.cpu.util", "CPU utilisation percentage.")
+}
