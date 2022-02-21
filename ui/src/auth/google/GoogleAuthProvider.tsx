import {
  createContext,
  ReactNode,
  useCallback,
  useEffect,
  useState,
} from "react";

import * as jose from "jose";

import { AuthContext, AuthState } from "../auth-context";
import { Google, useGoogleSdk } from "./use-google-sdk";
import { User } from "../user";
import { Backdrop, CircularProgress } from "@mui/material";

export const GoogleContext = createContext<Google.Id.Sdk>(null!);

export function GoogleAuthProvider(props: {
  children: ReactNode;
  clientId: string;
}): JSX.Element {
  const [token, setToken] = useState<string>();
  const sdk = useGoogleSdk(props.clientId, ({ credential }) =>
    setToken(credential)
  );

  const logout = useCallback(() => {
    sdk?.disableAutoSelect();
    setToken(undefined);
  }, [sdk]);

  const login = useCallback(() => {
    sdk?.prompt(console.log);
  }, [sdk]);

  useEffect(() => {
    if (!sdk) return;

    sdk.initialize({
      callback: ({ credential }) => setToken(credential),
      client_id: props.clientId,
      auto_select: true,
    });
  }, [sdk, props.clientId]);

  const authState: AuthState = token
    ? { isLoggedIn: true, user: parseUser(token), logout }
    : {
        isLoggedIn: false,
        login,
      };

  return sdk ? (
    <GoogleContext.Provider value={sdk}>
      <AuthContext.Provider value={authState}>
        {props.children}
      </AuthContext.Provider>
    </GoogleContext.Provider>
  ) : (
    <Backdrop
      sx={{ color: "#fff", zIndex: (theme) => theme.zIndex.drawer + 1 }}
      open={true}
    >
      <CircularProgress color="inherit" />
    </Backdrop>
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
