import React, { useEffect, useState } from "react";
import type { PerfDetails } from "@/services/perf.service";
import { getPerf } from "@/services/perf.service";

const PerfDashboard: React.FC = () => {
  const [perf, setPerf] = useState<PerfDetails | null>(null);

  useEffect(() => {
    let cancelled = false;
    const tick = async () => {
      const p = await getPerf();
      if (!cancelled) setPerf(p);
    };
    void tick();
    const id = window.setInterval(() => { void tick(); }, 2000);
    return () => {
      cancelled = true;
      window.clearInterval(id);
    };
  }, []);

  if (!perf) {
    return (
      <div className="rounded-xl p-4 bg-gradient-to-br from-muted/40 to-card border border-border text-sm text-muted-foreground">
        Performances indisponibles
      </div>
    );
  }

  const cpuChargeRaw = Number(perf.cpu_charge ?? 0);
  const cpuChargePct = Math.max(0, Math.min(100, Math.round(cpuChargeRaw)));
  const usedRam = Number(perf.used_ram ?? 0);
  const freeRam = Number(perf.dispo_ram ?? 0);
  const totalRam = usedRam + freeRam;
  const ramPct = totalRam > 0 ? Math.round((usedRam / totalRam) * 100) : 0;
  const cpuPower = Math.max(0, Number(perf.cpu_power ?? 0));
  const consumptionWh = cpuPower; // Puissance instantanée exprimée en Wh sur 1h

  const cpuBarColor = cpuChargePct < 50 ? "bg-green-500" : cpuChargePct < 80 ? "bg-amber-500" : "bg-red-600";
  const ramBarColor = ramPct < 60 ? "bg-blue-500" : ramPct < 85 ? "bg-amber-500" : "bg-red-600";

  return (
    <div className="rounded-xl p-4 bg-gradient-to-br from-orange-soft/40 to-cream border border-primary/30">
      <div className="text-sm font-semibold text-foreground mb-3">Système</div>

      <div className="space-y-4">
        {/* CPU Charge */}
        <div>
          <div className="flex items-center justify-between text-xs text-muted-foreground mb-1">
            <span>Charge CPU</span>
            <span className="font-medium text-foreground">{cpuChargePct}%</span>
          </div>
          <div className="h-2 w-full rounded-full bg-muted overflow-hidden">
            <div className={`h-full ${cpuBarColor}`} style={{ width: `${cpuChargePct}%` }} />
          </div>

        {/* Consommation instantanée (Wh) */}
        <div className="flex items-center justify-between text-xs">
          <span className="text-muted-foreground">Consommation</span>
          <span className="inline-flex items-center px-2 py-0.5 rounded-full bg-secondary/20 text-secondary-foreground font-medium">
            {consumptionWh.toFixed(1)} Wh
          </span>
        </div>
        </div>

        {/* CPU Power */}
        <div className="flex items-center justify-between text-xs">
          <span className="text-muted-foreground">Puissance CPU</span>
          <span className="inline-flex items-center px-2 py-0.5 rounded-full bg-primary/15 text-primary font-medium">
            {cpuPower.toFixed(1)} W
          </span>
        </div>

        {/* RAM */}
        <div>
          <div className="flex items-center justify-between text-xs text-muted-foreground mb-1">
            <span>Mémoire</span>
            <span className="font-medium text-foreground">
              {usedRam.toFixed(1)} / {totalRam.toFixed(1)} Go ({ramPct}%)
            </span>
          </div>
          <div className="h-2 w-full rounded-full bg-muted overflow-hidden">
            <div className={`h-full ${ramBarColor}`} style={{ width: `${ramPct}%` }} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default PerfDashboard;


