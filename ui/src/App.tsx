import logo from "./logo.svg";
import "./App.css";
import { useApi } from "./use-api";
import { useGoogleIdSdk } from "./use-google-id-sdk";
import { useEffect, useRef } from "react";
import { useAuth } from "./auth";

function App() {
  const apiConfig = useApi();
  const { setToken, isSignedIn, setSignoutCallback, signOut } = useAuth();

  if (!apiConfig) return <Loading />;

  if (!isSignedIn) {
    return (
      <GoogleLogin
        clientId={apiConfig.clientId}
        tokenCallback={setToken}
        registerSignoutCallback={setSignoutCallback}
      />
    );
  }

  return (
    <>
      <h1>Welcome!</h1>
      <button onClick={signOut}>Sign Out</button>
    </>
  );
}

export function Loading() {
  return <h1>Loading</h1>;
}

export function GoogleLogin(props: {
  clientId: string;
  tokenCallback: (token: string) => void;
  registerSignoutCallback: (cb: () => void) => void;
}) {
  const googleSdk = useGoogleIdSdk();
  const element = useRef<HTMLDivElement>(null);

  useEffect(() => {
    googleSdk?.initialize({
      callback: ({ credential }) => {
        props.tokenCallback(credential);
        props.registerSignoutCallback(() => googleSdk.disableAutoSelect());
      },
      client_id: props.clientId,
    });
  }, [googleSdk, props.tokenCallback, props.clientId]);

  useEffect(() => {
    if (element.current === null) return;
    googleSdk?.renderButton(element.current, { type: "standard" });
  }, [googleSdk]);

  return <div ref={element} />;
}

export default App;
