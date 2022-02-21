import { useState } from "react";
import * as jose from "jose";

export declare namespace Google.Id {
  interface Config {
    client_id: string;
    callback: (token: { credential: string }) => void;
    ux_mode?: "popup" | "redirect";
    auto_select?: boolean;
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

export function useGoogleSdk(
  clientId: string,
  tokenCallback: (token: { credential: string }) => void
): Google.Id.Sdk | undefined {
  const tries = useState(0); // Use this to trigger a refresh

  // This will retry every 50 ms until the library is loaded
  if (!window.google) {
    setTimeout(() => {
      tries[1]((tries) => tries + 1);
    }, 50);
    return undefined;
  }

  const sdk = window.google.accounts.id;
  sdk.initialize({
    callback: tokenCallback,
    client_id: clientId,
  });
  return sdk;
}
