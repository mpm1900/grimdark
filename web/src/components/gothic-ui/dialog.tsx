import * as React from 'react'

import { cn } from '#/lib/utils'
import {
  Dialog as DialogPrimitive,
  AlertDialog as AlertDialogPrimitive,
} from 'radix-ui'
import { DialogOverlay, DialogPortal } from '../ui/dialog'
import { GothicCloseButton } from './button'
import { AlertDialogOverlay, AlertDialogPortal } from '../ui/alert-dialog'

function GothicDialogContent({
  className,
  children,
  showCloseButton = true,
  ...props
}: React.ComponentProps<typeof DialogPrimitive.Content> & {
  showCloseButton?: boolean
}) {
  return (
    <DialogPortal data-slot="dialog-portal">
      <DialogOverlay />
      <DialogPrimitive.Content
        data-slot="dialog-content"
        className={cn(
          'fixed top-[50%] left-[50%] z-50 grid w-full max-w-[calc(100%-2rem)] translate-x-[-50%] translate-y-[-50%] gap-4 p-0 shadow-lg duration-200 outline-none data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95 sm:max-w-lg',
          'border-[10px] border-solid border-transparent bg-gradient-to-b from-neutral-950 to-[oklch(10%_0_0)] text-card-foreground',
          'bg-clip-border',
          "[border-image-source:url('/gothic/SkillFrameVert.png')]",
          '[border-image-slice:22_fill]',
          '[border-image-repeat:stretch]',
          'pt-8 font-serif',
          className
        )}
        {...props}
      >
        {children}
        {showCloseButton && (
          <DialogPrimitive.Close data-slot="dialog-close" asChild>
            <GothicCloseButton className="absolute z-10 -right-1 -top-1 shadow-[0px_0px_6px_1px_rgba(0,_0,_0,_0.5)]">
              <span className="sr-only">Close</span>
            </GothicCloseButton>
          </DialogPrimitive.Close>
        )}
      </DialogPrimitive.Content>
    </DialogPortal>
  )
}

function GothicAlertDialogContent({
  className,
  ...props
}: React.ComponentProps<typeof AlertDialogPrimitive.Content>) {
  return (
    <AlertDialogPortal>
      <AlertDialogOverlay />
      <AlertDialogPrimitive.Content
        data-slot="alert-dialog-content"
        className={cn(
          'fixed top-[50%] left-[50%] z-50 grid w-full max-w-[calc(100%-2rem)] translate-x-[-50%] translate-y-[-50%] gap-4 p-0 shadow-lg duration-200 outline-none data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95 sm:max-w-lg',
          'border-[10px] border-solid border-transparent bg-gradient-to-b from-neutral-950 to-[oklch(10.5%_0_0)] text-card-foreground',
          'bg-clip-border',
          "[border-image-source:url('/gothic/SkillFrameVert.png')]",
          '[border-image-slice:22_fill]',
          '[border-image-repeat:stretch]',
          'pt-8 font-serif',
          className
        )}
        {...props}
      />
    </AlertDialogPortal>
  )
}

function GothicDialogHeader({
  className,
  children,
  ...props
}: React.ComponentProps<'div'>) {
  return (
    <div
      data-slot="dialog-header"
      className={cn(
        'absolute -top-7 left-1/2 z-10 flex h-12 w-[calc(100%-2rem)] -translate-x-1/2 flex-col items-center justify-center gap-1 px-16 pt-3.5 pb-2 text-center text-card-foreground',
        'overflow-visible',
        className
      )}
      {...props}
    >
      <span
        aria-hidden
        className="absolute inset-y-0 left-16 right-16 -z-20 bg-[url('/gothic/TitleFrameNormal_Middle.png')] bg-center bg-no-repeat [background-size:100%_100%]"
      />
      <span
        aria-hidden
        className="absolute inset-y-0 left-0 -z-10 w-16 bg-[url('/gothic/TitleFrameNormal.png')] bg-left bg-no-repeat [background-size:auto_100%]"
      />
      <span
        aria-hidden
        className="absolute inset-y-0 right-0 -z-10 w-16 bg-[url('/gothic/TitleFrameNormal.png')] bg-right bg-no-repeat [background-size:auto_100%]"
      />
      {children}
    </div>
  )
}

function GothicDialogTitle({
  className,
  ...props
}: React.ComponentProps<typeof DialogPrimitive.Title>) {
  return (
    <DialogPrimitive.Title
      data-slot="dialog-title"
      className={cn(
        'font-cinzel-dec text-lg leading-none font-semibold text-foreground/60',
        '[text-shadow:1px_2px_0_var(--color-black)]',
        className
      )}
      {...props}
    />
  )
}

function GothicAlertDialogTitle({
  className,
  ...props
}: React.ComponentProps<typeof AlertDialogPrimitive.Title>) {
  return (
    <AlertDialogPrimitive.Title
      data-slot="alert-dialog-title"
      className={cn(
        'font-cinzel-dec text-lg leading-none font-semibold text-foreground/60',
        className
      )}
      {...props}
    />
  )
}

export {
  GothicDialogContent,
  GothicAlertDialogContent,
  GothicDialogHeader,
  GothicDialogTitle,
  GothicAlertDialogTitle,
}
