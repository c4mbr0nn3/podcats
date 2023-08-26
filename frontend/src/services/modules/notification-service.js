import { NotificationApi } from "@/api";
import { showSuccessSnackbar, showErrorSnackbar } from "@/utils/snackbar";

export async function getAll() {
  const { data } = await NotificationApi.getAll();
  return data;
}

export async function setAllRead() {
  let data = null;
  try {
    data = await NotificationApi.setAllRead();
    showSuccessSnackbar("All notifications marked as read");
  } catch {
    showErrorSnackbar("Error marking all notifications as read");
  }
  return data;
}

export async function setRead(id) {
  let data = null;
  try {
    data = await NotificationApi.setRead(id);
    showSuccessSnackbar("Notification marked as read");
  } catch {
    showErrorSnackbar("Error marking notification as read");
  }
  return data;
}
