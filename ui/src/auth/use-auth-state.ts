import { useCallback, useState } from "react";

export interface AuthState {
  isSignedIn: boolean;
  setToken(token?: string): void;
  signOut(): void;
  setSignoutCallback(cb: () => void): void;
}

export function useAuthState(): AuthState {
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
