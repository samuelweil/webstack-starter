import { ReactNode, useCallback, useEffect, useState } from "react";

import * as jose from "jose";

import { AuthContext, AuthState } from "../auth-context";
import { Google, useGoogleIdSdk } from "./use-google-id-sdk";
import { User } from "../user";

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
    ? { isLoggedIn: true, user: parseUser(token), logout }
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

function parseUser(token: string): User {
  const claims = jose.decodeJwt(token) as Google.Id.IdToken;
  return {
    name: claims.name,
    email: claims.email,
    picture: claims.picture,
  };
}
