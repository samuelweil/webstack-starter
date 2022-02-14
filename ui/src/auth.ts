import { useCallback, useState } from "react";

export function useAuth() {
  const [token, setToken] = useState<string>();
  const [signOutCallback, setSignoutCallback] = useState<() => void>();

  const signOut = useCallback(() => {
    setToken(undefined);
    if (signOutCallback) signOutCallback();
  }, [signOutCallback]);

  return {
    isSignedIn: !!token,
    setToken,
    signOut,
    setSignoutCallback,
  };
}
