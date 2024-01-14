import Link from 'next/link';
import { TweetWindow } from '@/components/tweet/Window';
import type { TTweet } from '@/types'

export default async function HomePage() {
  const res: Response = await fetch(
    `${process.env.API_URL}/tweets`,
    { cache: 'no-store' }
  );
  const tweets: TTweet[] = await res.json();
  return (
    <>
      <TweetWindow tweets={tweets} />
    </>
  );
};

