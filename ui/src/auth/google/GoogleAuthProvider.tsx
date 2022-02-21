import { ReactNode, useCallback, useEffect, useState } from "react";
import { AuthContext, AuthState } from "../auth-context";
import { useGoogleIdSdk } from "./use-google-id-sdk";

export function GoogleAuthProvider(props: {
  children: ReactNode;
  clientId: string;
}) {
  const googleApi = useGoogleIdSdk();
  const [token, setToken] = useState<string>();

  const logout = useCallback(() => {
    googleApi?.disableAutoSelect();
    setToken(undefined);
  }, [googleApi]);

  const login = useCallback(() => {
    googleApi?.prompt();
  }, [googleApi]);

  useEffect(() => {
    if (!googleApi) return;

    googleApi.initialize({
      callback: ({ credential }) => setToken(credential),
      client_id: props.clientId,
    });
  });

  const authState: AuthState = token
    ? { isLoggedIn: true, logout }
    : {
        isLoggedIn: false,
        login,
      };

  return (
    <AuthContext.Provider value={authState}>
      {props.children}
    </AuthContext.Provider>
  );
}
