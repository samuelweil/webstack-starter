import { useContext, useEffect, useState } from "react";
import { GoogleContext } from "./GoogleAuthProvider";

export function GoogleLoginButton() {
  const googleApi = useContext(GoogleContext);
  const [element, setElement] = useState<HTMLDivElement | null>(null);

  useEffect(() => {
    if (!element) return;

    googleApi.renderButton(element, { type: "standard" });
  }, [element, googleApi]);

  return <div ref={setElement} />;
}
