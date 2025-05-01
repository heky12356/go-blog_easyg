"use client";
import { usePathname } from "next/navigation";

export default function TitleSuffix() {
  const pathname = usePathname();

  let titleSuffix = "";
  if (pathname !== "/") {
    titleSuffix = pathname.split("/")[1];
  }

  return <div>{titleSuffix ? " | " + titleSuffix : ""}</div>;
}
