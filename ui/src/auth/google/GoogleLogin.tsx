import { useGoogleIdSdk } from "./use-google-id-sdk";
import { useEffect, useRef } from "react";
import { useAuth } from "../use-auth";

export function GoogleLogin(props: { clientId: string }) {
  const googleSdk = useGoogleIdSdk();
  const element = useRef<HTMLDivElement>(null);
  const { setToken, setSignoutCallback } = useAuth();
  useEffect(() => {
    googleSdk?.initialize({
      callback: ({ credential }) => {
        setToken(credential);
        setSignoutCallback(() => googleSdk.disableAutoSelect());
      },
      client_id: props.clientId,
    });
  }, [googleSdk, setToken, setSignoutCallback, props.clientId]);

  useEffect(() => {
    if (element.current === null) return;
    googleSdk?.renderButton(element.current, { type: "standard" });
  }, [googleSdk]);

  return <div ref={element} />;
}
