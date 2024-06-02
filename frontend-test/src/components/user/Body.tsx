import { TTweet } from "@/types"
import { TweetContainer } from "@/components/tweet/Container"

export const UserBody = ({ tweets }: { tweets: TTweet[] }) => {
  return (
    <TweetContainer tweets={tweets} />
  )
}
