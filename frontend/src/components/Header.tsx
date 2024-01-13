import Link from "next/link";

export default function Header() {
  return (
    <>
      <header>
        <div>==========</div>
        <div>Header!!!</div>
        <div><Link href={'/home'}>home</Link></div>
        <div><Link href="/tweet/create">create</Link></div>
        <div><Link href={'/test'}>test</Link></div>
        <div>==========</div>
      </header>
    </>
  );
};

