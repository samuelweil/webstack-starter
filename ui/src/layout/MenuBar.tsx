import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import { useAuth, User } from "../auth";
import { Avatar, Menu, MenuItem, Tooltip } from "@mui/material";
import { useState } from "react";
import { GoogleLoginButton } from "../auth/google";

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
            <UserMenu user={auth.user} logOut={() => auth.logout()} />
          ) : (
            <GoogleLoginButton />
          )}
        </Toolbar>
      </AppBar>
    </Box>
  );
}

function UserMenu(props: { user: User; logOut: () => void }) {
  const [open, setOpen] = useState(false);
  const [element, setElement] = useState<HTMLButtonElement | null>(null);

  return (
    <>
      <Tooltip title="Open settings">
        <IconButton
          ref={setElement}
          sx={{ p: 0 }}
          onClick={() => setOpen((open) => !open)}
        >
          <Avatar alt="Remy Sharp" src={props.user.picture} />
        </IconButton>
      </Tooltip>
      <Menu
        anchorEl={element}
        anchorOrigin={{
          vertical: "top",
          horizontal: "right",
        }}
        sx={{ mt: "45px" }}
        keepMounted
        transformOrigin={{
          vertical: "top",
          horizontal: "right",
        }}
        open={open}
        onClose={() => setOpen((open) => !open)}
      >
        <MenuItem key="1" onClick={props.logOut}>
          <Typography textAlign="center">Logout</Typography>
        </MenuItem>
      </Menu>
    </>
  );
}
