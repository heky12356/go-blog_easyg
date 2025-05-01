import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import 'bootstrap/dist/css/bootstrap.min.css';
import { Container } from "react-bootstrap";
import Link from "next/link";
const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "EasyG",
  description: "",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <head>
      <script src="https://cdn.jsdelivr.net/npm/react/umd/react.production.min.js" ></script>
      </head>
      <body className={`${geistSans.variable} ${geistMono.variable}`}>
        <div>
      <Container>
        <div className="d-flex">
          <h1 className="pt-4">
            <Link href={"/"} className="text-decoration-none text-black">
              Blog
            </Link>
          </h1>
          {/* <div className="align-self-end pb-2">
            {titleSuffix ? " | " + titleSuffix : ""}
          </div> */}
        </div>
        <hr className="mb-5" />
        <Container className="container-height">
        {children}
        </Container>
        {/* <Footer /> */}
      </Container>
    </div>
      </body>
    </html>
  );
}
