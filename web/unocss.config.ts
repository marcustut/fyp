import { defineConfig, presetIcons, presetUno } from "unocss";
import { presetTypography } from "unocss-preset-typography";

export default defineConfig({
  shortcuts: [],
  presets: [presetUno(), presetIcons(), presetTypography()],
});
