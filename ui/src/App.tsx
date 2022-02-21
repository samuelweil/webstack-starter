import { useEffect } from "react";

import { Container } from "@mui/material";

import { useApi } from "./use-api";
import { GoogleAuthProvider, useAuth } from "./auth";
import { MenuBar } from "./layout/MenuBar";

function App() {
  const apiConfig = useApi();

  useEffect(() => {
    if (apiConfig) console.log("Api loaded");
  }, [apiConfig]);

  if (!apiConfig) return <Loading />;

  return (
    <GoogleAuthProvider clientId={apiConfig.clientId}>
      <Home />
    </GoogleAuthProvider>
  );
}

function Home() {
  const authState = useAuth();

  return (
    <>
      <MenuBar />
      <Container>
        <h1>{authState.isLoggedIn ? "Hello!" : "Please login"}</h1>
      </Container>
    </>
  );
}

function Loading() {
  return <h1>Loading</h1>;
}

export default App;
