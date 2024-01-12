import Link from 'next/link';
import { TweetWindow } from '@/components/TweetWindow';
import type { TTweet } from '@/types'

export default async function Home() {
  const res: Response = await fetch(
    `${process.env.API_URL}/tweets`,
    { cache: 'no-store' }
  );
  const tweets: TTweet[] = await res.json();
  return (
    <>
      <TweetWindow tweets={tweets} />
      <Link href={'/tweet/create'}>
        create
      </Link>
    </>
  );
};

