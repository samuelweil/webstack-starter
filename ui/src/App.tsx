import logo from "./logo.svg";
import "./App.css";
import { useApi } from "./use-api";
import { useGoogleIdSdk } from "./use-google-id-sdk";
import { useEffect, useRef } from "react";
import { useAuth } from "./auth";

function App() {
  const apiConfig = useApi();
  const { setToken, isSignedIn } = useAuth();

  if (!apiConfig) return <Loading />;

  if (!isSignedIn) {
    return (
      <GoogleLogin clientId={apiConfig.clientId} tokenCallback={setToken} />
    );
  }

  return <h1>Welcome!</h1>;
}

export function Loading() {
  return <h1>Loading</h1>;
}

export function GoogleLogin(props: {
  clientId: string;
  tokenCallback: (token: string) => void;
}) {
  const googleSdk = useGoogleIdSdk();
  const element = useRef<HTMLDivElement>(null);

  useEffect(() => {
    googleSdk?.initialize({
      callback: ({ credential }) => props.tokenCallback(credential),
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
