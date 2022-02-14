import "./App.css";
import { useApi } from "./use-api";
import { AuthProvider, GoogleLogin, useAuth } from "./auth";

function App() {
  const apiConfig = useApi();
  if (!apiConfig) return <Loading />;

  return (
    <AuthProvider>
      <Home clientId={apiConfig.clientId} />
    </AuthProvider>
  );
}

function Home(props: { clientId: string }) {
  const { isSignedIn, signOut } = useAuth();

  return isSignedIn ? (
    <>
      <h1>Welcome!</h1>
      <button onClick={signOut}>Sign Out</button>
    </>
  ) : (
    <GoogleLogin clientId={props.clientId} />
  );
}

function Loading() {
  return <h1>Loading</h1>;
}

export default App;
