import { TweetContainer } from '@/components/tweet/Container';
import type { TTweet } from '@/types'

export default async function HomePage() {
  const res: Response = await fetch(
    `${process.env.API_URL}/tweets`,
    { cache: 'no-store' }
  );
  const tweets: TTweet[] = await res.json();
  return (
    <>
      <TweetContainer tweets={tweets} />
    </>
  );
};

