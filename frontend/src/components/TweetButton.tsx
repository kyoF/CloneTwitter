'use client';

import { useRouter } from 'next/navigation';

export const TweetButton = ({ userId }: { userId: string }) => {
  const router = useRouter()
  const tweet = () => {
    router.push('/tweet/create');
  };
  return (
    <button onClick={() => tweet()}>Let's Tweet!!</button>
  );
}
