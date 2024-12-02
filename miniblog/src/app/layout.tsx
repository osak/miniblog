import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
    title: 'Miniblog',
    description: 'A minimal blog',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body>
      <div id="header">
        <h1>Miniblog</h1>
      </div>
      <div id="main">
        {children}
      </div>
      </body>
    </html>
  );
}
