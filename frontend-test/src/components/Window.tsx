import Sidebar from '@/components/Sidebar';
import Search from '@/components/tweet/Search';

export default function Window({ children }: { children: React.ReactNode }) {
  return (
    <div>
      <div>
        <Sidebar />
      </div>
      <div>
        {children}
      </div>
      <div>
        <Search />
      </div>
    </div>
  )
}
