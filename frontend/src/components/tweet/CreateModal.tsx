'use client';

export const TweetCreateModal = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="fixed z-10 inset-0 overflow-y-auto bg-black bg-opacity-40 flex items-center justify-center">
      <div className="bg-white rounded-lg p-6 w-1/3">
        <div className="flex justify-between items-center mb-6">
          {children}
        </div>
      </div>
    </div>
  );
}
