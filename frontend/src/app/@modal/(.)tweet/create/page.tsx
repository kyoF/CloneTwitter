'use client';

import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { TweetCreateModal } from '@/components/tweet/CreateModal';

export default function TweetCreateModalPage() {
  const router = useRouter();
  const [tweetText, setTweetText] = useState<string>('');
  const [characterCount, setCharacterCount] = useState<integer>(0);

  const handleTweetTextChange = (e) => {
    const text = e.target.value;
    setTweetText(text);
    setCharacterCount(text.length);
  };

  const handleTweetSubmit = async (e) => {
    e.preventDefault()

    const res: Response = await fetch('/api/tweet', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ userId: 'fujiki', text: tweetText })
    })
    setTweetText('');
    setCharacterCount(0);
    router.back()
  };

  const back = () => {
    router.back()
  }

  return (
    <TweetCreateModal>
      <textarea
        placeholder="What's happeing?"
        value={tweetText}
        onChange={handleTweetTextChange}
        maxLength={280}
      />
      <div>
        <span>{characterCount}/280</span>
        <button
          onClick={handleTweetSubmit}
        >
          tweet
        </button>
      </div>
      <button onClick={back}>back</button>
    </TweetCreateModal>
  );
}

