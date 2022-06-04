import "emoji-picker-element";

import ClipboardJS from "clipboard";
import { useEffect } from "react";

function App() {
  useEffect(() => {
    document
      .querySelector("emoji-picker")
      .addEventListener("emoji-click", (event) => {
        ClipboardJS.copy(event.detail.unicode);
      });
  }, []);

  return <emoji-picker></emoji-picker>;
}

export default App;
