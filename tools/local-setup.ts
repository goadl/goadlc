import {
  packages,
  forPlatform,
  getHostPlatform,
  installTo,
  MultiPlatform,
} 
from "https://deno.land/x/adllang_localsetup@v0.9/mod.ts";

const platform = getHostPlatform();

function withPlatform<T>(multi: MultiPlatform<T>) {
  return forPlatform(multi, platform);
}

const ADL = withPlatform(packages.adl("1.1.12"));

export async function main() {
  if (Deno.args.length != 1) {
    console.error("Usage: local-setup LOCALDIR");
    Deno.exit(1);
  }
  const localdir = Deno.args[0];


  const installs = [
    ADL,
  ];

  await installTo(installs, localdir);
}

main()
  .catch((err) => {
    console.error("error in main", err);
  });
