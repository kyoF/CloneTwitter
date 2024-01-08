import { Tweet } from '@/components/Tweet'
import type { TTweet } from '@/types'

export default async function Home() {
  const res: Response = await fetch(`${process.env.API_URL}/tweets`, { cache: 'no-store' });
  const tweets: TTweet[] = await res.json();
  return (
    <>
      <h1>Home</h1>
      {tweets.map((tweet: TTweet) => (
        <div key={tweet.tweetId}>
          <Tweet {...tweet} />
        </div>
      ))}
    </>
  );
};

