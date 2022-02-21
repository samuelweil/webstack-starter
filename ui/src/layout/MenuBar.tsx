import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import { useAuth, User } from "../auth";
import { Avatar, Tooltip } from "@mui/material";

export function MenuBar() {
  const auth = useAuth();

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <IconButton
            size="large"
            edge="start"
            color="inherit"
            aria-label="menu"
            sx={{ mr: 2 }}
          >
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Workflow
          </Typography>
          {auth.isLoggedIn ? (
            <UserMenu user={auth.user} />
          ) : (
            <Button color="inherit" onClick={() => auth.login()}>
              Login
            </Button>
          )}
        </Toolbar>
      </AppBar>
    </Box>
  );
}

function UserMenu(props: { user: User }) {
  return (
    <Tooltip title="Open settings">
      <IconButton sx={{ p: 0 }}>
        <Avatar alt="Remy Sharp" src={props.user.picture} />
      </IconButton>
    </Tooltip>
  );
}
