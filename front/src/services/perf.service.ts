export type PerfDetails = {
  cpu_charge: number;
  cpu_temp: number;
  cpu_power: number;
  used_ram: number;
  dispo_ram: number;
};

const API_URL = "http://localhost:8080/api/v1/perf";

type PerfApiResponse = {
  status: string;
  message: string;
  data?: PerfDetails;
};

export async function getPerf(): Promise<PerfDetails | null> {
  try {
    const res = await fetch(API_URL, { headers: { Accept: "application/json" } });
    if (!res.ok) {
      console.error("Perf API status:", res.status);
      return null;
    }
    const contentType = res.headers.get("content-type") || "";
    if (!contentType.includes("application/json")) {
      const text = await res.text();
      console.error("Perf non-JSON:", text.slice(0, 200));
      return null;
    }
    const json = (await res.json()) as PerfApiResponse | PerfDetails;
    // Support both {status,message,data} and raw details for safety
    const details: PerfDetails | undefined = (json as PerfApiResponse).data
      ? (json as PerfApiResponse).data
      : (json as PerfDetails);
    if (!details) return null;
    return details;
  } catch (e) {
    console.error("Perf fetch error:", e);
    return null;
  }
}


