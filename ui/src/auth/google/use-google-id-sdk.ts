import { useEffect, useState } from "react";

declare namespace Google.Id {
  interface Config {
    client_id: string;
    callback: (token: { credential: string }) => void;
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
