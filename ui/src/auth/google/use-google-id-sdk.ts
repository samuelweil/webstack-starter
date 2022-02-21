import { useEffect, useState } from "react";
import * as jose from "jose";

// "iss": "https://accounts.google.com", // The JWT's issuer
// "nbf":  161803398874,
// "aud": "314159265-pi.apps.googleusercontent.com", // Your server's client ID
// "sub": "3141592653589793238", // The unique ID of the user's Google Account
// "hd": "gmail.com", // If present, the host domain of the user's GSuite email address
// "email": "elisa.g.beckett@gmail.com", // The user's email address
// "email_verified": true, // true, if Google has verified the email address
// "azp": "314159265-pi.apps.googleusercontent.com",
// "name": "Elisa Beckett",
//                           // If present, a URL to user's profile picture
// "picture": "https://lh3.googleusercontent.com/a-/e2718281828459045235360uler",
// "given_name": "Elisa",
// "family_name": "Beckett",
// "iat": 1596474000, // Unix timestamp of the assertion's creation time
// "exp": 1596477600, // Unix timestamp of the assertion's expiration time
// "jti": "abc161803398874def"

export declare namespace Google.Id {
  interface Config {
    client_id: string;
    callback: (token: { credential: string }) => void;
  }

  export interface IdToken extends jose.JWTPayload {
    email: string;
    name: string;
    picture: string;
  }

  interface PromptMoment {
    isDismissedMoment(): boolean;
    getDismissedReason():
      | "credential_returned"
      | "cancel_called"
      | "flow_restarted";
  }

  interface GsiButtonConfiguration {
    type: "icon" | "standard";
  }

  export interface Sdk {
    initialize(idConfig: Config): void;
    prompt(listener?: (moment: PromptMoment) => void): void;
    renderButton(htmlEl: HTMLElement, config: GsiButtonConfiguration): void;
    disableAutoSelect(): void;
  }
}

declare global {
  interface Window {
    google?: { accounts: { id: Google.Id.Sdk } };
  }
}

export function useGoogleIdSdk(): Google.Id.Sdk | undefined {
  const [sdk, setSdk] = useState<Google.Id.Sdk>();

  useEffect(() => {
    const gSdk = window.google?.accounts.id;
    if (gSdk !== undefined) {
      setSdk(gSdk);
    }
  }, []);

  return sdk;
}
