'use client'

import Link from 'next/link';
import { useState } from 'react';
import { CiUser, CiHeart } from "react-icons/ci";

import type { TTweet } from '@/types';

export const TweetCard = ({ ...tweet }: TTweet) => {
  const { userId, text, likesCount } = tweet;

  const [like, setLike] = useState<number>(likesCount);
  const [isLike, setIsLike] = useState<Boolean>(false);

  const toggleLike = () => {
    if (isLike) {
      setLike(like - 1);
    } else {
      setLike(like + 1);
    }
    setIsLike(!isLike)
  }

  return (
    <>
      <CiUser />
      <Link href={`/user/${userId}`}>
        @{userId}
      </Link>
      <div>{text}</div>
      <button onClick={toggleLike}>
        <div>
          <CiHeart />
        </div>
      </button>
      {like}
    </>
  );
}

