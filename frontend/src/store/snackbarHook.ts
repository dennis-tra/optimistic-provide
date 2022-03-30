import React, { useEffect } from "react";
import { SnackbarKey, useSnackbar } from "notistack";
import { actions } from "./snackbarSlice";
import { useAppDispatch, useAppSelector } from "./config";
import { Button } from "@mui/material";
import CloseIcon from "@mui/icons-material/Close";

let displayed: SnackbarKey[] = [];

const useNotifier = () => {
  const dispatch = useAppDispatch();
  const notifications = useAppSelector((store) => store.snackbar.notifications);
  const { enqueueSnackbar, closeSnackbar } = useSnackbar();

  const storeDisplayed = (key: SnackbarKey) => {
    displayed = [...displayed, key];
  };

  const removeDisplayed = (id: SnackbarKey) => {
    displayed = [...displayed.filter((key) => id !== key)];
  };

  useEffect(() => {
    notifications.forEach(({ key, message, action, variant, onClose, dismissed = false }) => {
      if (dismissed) {
        closeSnackbar(key);
        return;
      }

      if (displayed.includes(key)) return;

      const fallbackAction = (key: SnackbarKey) =>
        React.createElement(
          Button,
          { onClick: () => dispatch(actions.dismissNotification(key)) },
          React.createElement(CloseIcon)
        );

      enqueueSnackbar(message, {
        key,
        variant,
        action: action || fallbackAction,
        onClose: (event, reason, myKey) => {
          if (onClose) {
            onClose(event, reason, myKey);
          }
        },
        onExited: (event, myKey) => {
          dispatch(actions.removeNotification(myKey));
          removeDisplayed(myKey);
        },
      });

      storeDisplayed(key);
    });
  }, [notifications, closeSnackbar, enqueueSnackbar, dispatch]);
};

export default useNotifier;
