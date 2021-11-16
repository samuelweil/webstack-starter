import { useEffect, useState } from "react";
import axios from "axios";

interface Info {
  version: string;
}

export function useApi(): Info | null {
  const [info, setInfo] = useState<Info | null>(null);
  useEffect(() => {
    setTimeout(async () => {
      const { data: info } = await axios.get<Info>("/api");
      setInfo(info);
    }, 2000);
  }, []);

  return info;
}
