import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";

import confetti from "canvas-confetti";
import { useEffect } from "react";

const Home: NextPage = () => {
  useEffect(() => {
    confetti();
  }, []);

  return (
    <div className="h-100vh bg-dark-400">
      <div className="i-mdi:alarm text-3xl text-dark-500" />
    </div>
  );
};

export default Home;
