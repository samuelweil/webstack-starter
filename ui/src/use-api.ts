import { useEffect, useState } from "react";
import axios from "axios";

interface Config {
  clientId: string;
}

export function useApi(): Config | undefined {
  const [config, setConfig] = useState<Config>();

  useEffect(() => {
    loadConfig().then(setConfig);
  }, []);

  return config;
}

export async function loadConfig(): Promise<Config> {
  const { data: config } = await axios.get<Config>("/api/config");
  return config;
}
