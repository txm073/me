const toggle = document.getElementById("theme-toggle");
const savedTheme = localStorage.getItem("theme") || "light";

document.documentElement.dataset.theme = savedTheme;
toggle.checked = savedTheme === "dark";

toggle.addEventListener("click", () => {
  const theme = toggle.checked ? "dark" : "light";
  document.documentElement.dataset.theme = theme;
  localStorage.setItem("theme", theme);
});
