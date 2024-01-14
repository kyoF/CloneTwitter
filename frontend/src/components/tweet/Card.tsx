'use client'

import Link from 'next/link';
import { useState } from 'react';
import { HeartIcon, PersonIcon } from '@radix-ui/react-icons';
import { css } from '@/../styled-system/css';

import type { TTweet } from '@/types';

export const TweetCard = ({...tweet}: TTweet) => {
  const { userId, text, likesCount } = tweet;

  const [like, setLike] = useState<number>(likesCount);
  const [isLike, setIsLike] = useState<Boolean>(false);

  const toggleLike = () => {
    if (isLike) {
      setLike(like-1);
    } else {
      setLike(like+1);
    }
    setIsLike(!isLike)
  }

  return(
    <>
      <PersonIcon className={css({
        width: "5%",
        height: "5%",
      })} />
      <Link href={`/user/${userId}`}>
        @{userId}
      </Link>
      <div>{text}</div>
      <button onClick={toggleLike}>
        <div className={css({
          color: isLike ? "red" : "black",
        })}>
          <HeartIcon />
        </div>
      </button>
      {like}
    </>
  );
}

