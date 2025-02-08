export const dynamicTitle = () => {
  const path = window.location.pathname;
  const lastSegment = path.substring(path.lastIndexOf("/") + 1);
  if (lastSegment === "") {
    document.title = "Chat Cuy";
    return;
  }
  const formattedPath = lastSegment
    .replace(/-/g, " ")
    .split(" ")
    .map((word) => {
      return word.charAt(0).toUpperCase() + word.slice(1);
    })
    .join(" ");
  document.title = `Chat Cuy | ${formattedPath}`;
};
