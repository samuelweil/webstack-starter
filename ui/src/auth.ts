import { useState } from "react";

export function useAuth() {
  const [token, setToken] = useState<string>();

  return {
    isSignedIn: !!token,
    setToken,
  };
}
