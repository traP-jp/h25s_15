import type { GameEvent } from '@/lib/type'
import { onMounted, onUnmounted } from 'vue'

export const useGameEvent = (wsUrl: string, onEvent: (event: GameEvent) => void) => {
  const ws = new WebSocket(wsUrl)

  const onMessage = (event: MessageEvent<string>) => {
    try {
      const gameEvent = JSON.parse(event.data) as GameEvent

      onEvent(gameEvent)
    } catch (err) {
      console.warn('Failed to handle WebSocket event:', err)
    }
  }

  onMounted(() => ws.addEventListener('message', onMessage))
  onUnmounted(() => ws.removeEventListener('message', onMessage))
}
