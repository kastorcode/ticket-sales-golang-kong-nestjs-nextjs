"use client"

import { selectSpotAction, unSelectSpotAction } from "~/actions"

interface SpotSeatProps {
  spotId    : string
  spotLabel : string
  eventId   : string
  selected  : boolean
  disabled  : boolean
}

export function SpotSeat ({ spotId, spotLabel, eventId, selected, disabled } : SpotSeatProps) {
  return (
    <div className="flex">
      <input
        type="checkbox"
        name={`spots`}
        id={`spot-${spotId}`}
        className="peer hidden"
        value={spotId}
        disabled={disabled}
        defaultChecked={selected}
        onChange={async event => {
          event.target.checked ?
            await selectSpotAction(eventId, spotId) :
            await unSelectSpotAction(spotId)
        }}
      />
      <label
        htmlFor={`spot-${spotId}`}
        className="m-1 h-6 w-6 cursor-pointer select-none
          rounded-full bg-available py-1 text-center
          text-[10px] text-black
          peer-checked:bg-chosen
          peer-disabled:cursor-default
          peer-disabled:bg-occupied
          "
      >
        {spotLabel}
      </label>
    </div>
  )
}