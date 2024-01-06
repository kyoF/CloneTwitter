"use client"

import { useEffect } from "react";

export default function Home({ data }) {
  useEffect(() => {},[])
  return (
    <>
      <p>Home</p>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </>
  );
};

export async function getServerSideProps() {
  const endpoint: string = "http://localhost:8888/tweets";
  const res: Response = await fetch(endpoint);
  const data = await res.json();
  return {
    props: {
      data,
    },
  };
};

