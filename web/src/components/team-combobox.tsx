import { useQuery } from '@tanstack/react-query'
import {
  Combobox,
  ComboboxContent,
  ComboboxEmpty,
  ComboboxInput,
  ComboboxItem,
  ComboboxList,
  ComboboxTrigger,
} from './ui/combobox'
import { teamsQuery } from '#/lib/queries/get-teams'
import type { Team } from '#/lib/game/team'
import { useDeleteTeam } from '#/lib/mutations/delete-team'
import { Trash2 } from 'lucide-react'
import { Button } from './ui/button'

function TeamCombobox({
  onSelect,
  ...props
}: React.ComponentProps<typeof ComboboxTrigger> & {
  onSelect: (team: Team) => void
}) {
  const query = useQuery(teamsQuery)
  const delete_team = useDeleteTeam()

  return (
    <Combobox
      items={query.data}
      onValueChange={(team_ID: string | null) => {
        const team = query.data?.find((t) => t.ID === team_ID)
        if (team) onSelect(team)
      }}
    >
      <ComboboxTrigger {...props} />
      <ComboboxContent className="min-w-40">
        <ComboboxInput showTrigger={false} placeholder="Search Teams" />
        <ComboboxEmpty>No teams found.</ComboboxEmpty>
        <ComboboxList>
          {(team: Team) => (
            <ComboboxItem
              value={team.ID}
              className="min-w-0"
              showIndicator={false}
              action={
                <Button
                  variant="ghost"
                  size="icon-xs"
                  className="cursor-pointer hover:underline"
                  aria-label={`Delete ${team.config.name}`}
                  disabled={delete_team.isPending}
                  onPointerDown={(e) => {
                    e.preventDefault()
                    e.stopPropagation()
                  }}
                  onClick={(e) => {
                    e.preventDefault()
                    e.stopPropagation()
                    if (team.ID) {
                      delete_team.mutate(team.ID)
                    }
                  }}
                >
                  <Trash2 />
                </Button>
              }
            >
              <span className="min-w-0 flex-1 truncate">
                {team.config.name}
              </span>
            </ComboboxItem>
          )}
        </ComboboxList>
      </ComboboxContent>
    </Combobox>
  )
}

export { TeamCombobox }
