import {createRoot} from "react-dom/client";
import {PostEditor} from "./PostEditor/PostEditor";
import React from "react";

document.addEventListener('DOMContentLoaded', () => {
    const dom = document.getElementById("main")!!;
    const root = createRoot(dom);
    root.render(<PostEditor />);
});