'use client'

import Link from 'next/link';
import { useState } from 'react';

import type { TTweet } from '@/types';

export const TweetCard = ({ ...tweet }: TTweet) => {
  const { userId, name, tweetId, text, likesCount } = tweet;

  const [like, setLike] = useState<number>(likesCount);
  const [isLike, setIsLike] = useState<Boolean>(false);

  const toggleLike = async (e) => {
    e.preventDefault()

    try {
      const res: Response = await fetch('/api/like', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ userId: 'fujiki', tweetId: tweetId })
      });

      if (!res.ok) {
        throw new Error(`API request failed with status ${res.status}`);
      }

      const data = await res.json()
      setLike(data.likesCount);
      setIsLike(!isLike);
    } catch (error) {
      console.error('Error submitting tweet:', error);
    }
  };

  return (
    <>
      <div>user icon</div>
      <div>
        <Link href={`/user/${userId}`}>
          @{userId}
        </Link>
        {name}
      </div>
      <div>{text}</div>
      <button onClick={toggleLike}>
        <div>
          heart icon
        </div>
      </button>
      {like}
    </>
  );
}

