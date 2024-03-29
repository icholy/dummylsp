package main

import (
	"context"
	"errors"

	"go.lsp.dev/protocol"
)

var ErrNotImplemented = errors.New("lsp: not implemented")

type DefaultServer struct{}

// CodeAction implements protocol.Server.
func (*DefaultServer) CodeAction(ctx context.Context, params *protocol.CodeActionParams) (result []protocol.CodeAction, err error) {
	return nil, ErrNotImplemented
}

// CodeLens implements protocol.Server.
func (*DefaultServer) CodeLens(ctx context.Context, params *protocol.CodeLensParams) (result []protocol.CodeLens, err error) {
	return nil, ErrNotImplemented
}

// CodeLensRefresh implements protocol.Server.
func (*DefaultServer) CodeLensRefresh(ctx context.Context) (err error) {
	return ErrNotImplemented
}

// CodeLensResolve implements protocol.Server.
func (*DefaultServer) CodeLensResolve(ctx context.Context, params *protocol.CodeLens) (result *protocol.CodeLens, err error) {
	return nil, ErrNotImplemented
}

// ColorPresentation implements protocol.Server.
func (*DefaultServer) ColorPresentation(ctx context.Context, params *protocol.ColorPresentationParams) (result []protocol.ColorPresentation, err error) {
	return nil, ErrNotImplemented
}

// Completion implements protocol.Server.
func (*DefaultServer) Completion(ctx context.Context, params *protocol.CompletionParams) (result *protocol.CompletionList, err error) {
	return nil, ErrNotImplemented
}

// CompletionResolve implements protocol.Server.
func (*DefaultServer) CompletionResolve(ctx context.Context, params *protocol.CompletionItem) (result *protocol.CompletionItem, err error) {
	return nil, ErrNotImplemented
}

// Declaration implements protocol.Server.
func (*DefaultServer) Declaration(ctx context.Context, params *protocol.DeclarationParams) (result []protocol.Location, err error) {
	return nil, ErrNotImplemented
}

// Definition implements protocol.Server.
func (*DefaultServer) Definition(ctx context.Context, params *protocol.DefinitionParams) (result []protocol.Location, err error) {
	return nil, ErrNotImplemented
}

// DidChange implements protocol.Server.
func (*DefaultServer) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (err error) {
	return ErrNotImplemented
}

// DidChangeConfiguration implements protocol.Server.
func (*DefaultServer) DidChangeConfiguration(ctx context.Context, params *protocol.DidChangeConfigurationParams) (err error) {
	return ErrNotImplemented
}

// DidChangeWatchedFiles implements protocol.Server.
func (*DefaultServer) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) (err error) {
	return ErrNotImplemented
}

// DidChangeWorkspaceFolders implements protocol.Server.
func (*DefaultServer) DidChangeWorkspaceFolders(ctx context.Context, params *protocol.DidChangeWorkspaceFoldersParams) (err error) {
	return ErrNotImplemented
}

// DidClose implements protocol.Server.
func (*DefaultServer) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
	return ErrNotImplemented
}

// DidCreateFiles implements protocol.Server.
func (*DefaultServer) DidCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (err error) {
	return ErrNotImplemented
}

// DidDeleteFiles implements protocol.Server.
func (*DefaultServer) DidDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (err error) {
	return ErrNotImplemented
}

// DidOpen implements protocol.Server.
func (*DefaultServer) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	return ErrNotImplemented
}

// DidRenameFiles implements protocol.Server.
func (*DefaultServer) DidRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (err error) {
	return ErrNotImplemented
}

// DidSave implements protocol.Server.
func (*DefaultServer) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	return ErrNotImplemented
}

// DocumentColor implements protocol.Server.
func (*DefaultServer) DocumentColor(ctx context.Context, params *protocol.DocumentColorParams) (result []protocol.ColorInformation, err error) {
	return nil, ErrNotImplemented
}

// DocumentHighlight implements protocol.Server.
func (*DefaultServer) DocumentHighlight(ctx context.Context, params *protocol.DocumentHighlightParams) (result []protocol.DocumentHighlight, err error) {
	return nil, ErrNotImplemented
}

// DocumentLink implements protocol.Server.
func (*DefaultServer) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) (result []protocol.DocumentLink, err error) {
	return nil, ErrNotImplemented
}

// DocumentLinkResolve implements protocol.Server.
func (*DefaultServer) DocumentLinkResolve(ctx context.Context, params *protocol.DocumentLink) (result *protocol.DocumentLink, err error) {
	return nil, ErrNotImplemented
}

// DocumentSymbol implements protocol.Server.
func (*DefaultServer) DocumentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) (result []interface{}, err error) {
	return nil, ErrNotImplemented
}

// ExecuteCommand implements protocol.Server.
func (*DefaultServer) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (result interface{}, err error) {
	return nil, ErrNotImplemented
}

// FoldingRanges implements protocol.Server.
func (*DefaultServer) FoldingRanges(ctx context.Context, params *protocol.FoldingRangeParams) (result []protocol.FoldingRange, err error) {
	return nil, ErrNotImplemented
}

// Formatting implements protocol.Server.
func (*DefaultServer) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) (result []protocol.TextEdit, err error) {
	return nil, ErrNotImplemented
}

// Hover implements protocol.Server.
func (*DefaultServer) Hover(ctx context.Context, params *protocol.HoverParams) (result *protocol.Hover, err error) {
	return nil, ErrNotImplemented
}

// Implementation implements protocol.Server.
func (*DefaultServer) Implementation(ctx context.Context, params *protocol.ImplementationParams) (result []protocol.Location, err error) {
	return nil, ErrNotImplemented
}

// IncomingCalls implements protocol.Server.
func (*DefaultServer) IncomingCalls(ctx context.Context, params *protocol.CallHierarchyIncomingCallsParams) (result []protocol.CallHierarchyIncomingCall, err error) {
	return nil, ErrNotImplemented
}

// LinkedEditingRange implements protocol.Server.
func (*DefaultServer) LinkedEditingRange(ctx context.Context, params *protocol.LinkedEditingRangeParams) (result *protocol.LinkedEditingRanges, err error) {
	return nil, ErrNotImplemented
}

// LogTrace implements protocol.Server.
func (*DefaultServer) LogTrace(ctx context.Context, params *protocol.LogTraceParams) (err error) {
	return ErrNotImplemented
}

// Moniker implements protocol.Server.
func (*DefaultServer) Moniker(ctx context.Context, params *protocol.MonikerParams) (result []protocol.Moniker, err error) {
	return nil, ErrNotImplemented
}

// OnTypeFormatting implements protocol.Server.
func (*DefaultServer) OnTypeFormatting(ctx context.Context, params *protocol.DocumentOnTypeFormattingParams) (result []protocol.TextEdit, err error) {
	return nil, ErrNotImplemented
}

// OutgoingCalls implements protocol.Server.
func (*DefaultServer) OutgoingCalls(ctx context.Context, params *protocol.CallHierarchyOutgoingCallsParams) (result []protocol.CallHierarchyOutgoingCall, err error) {
	return nil, ErrNotImplemented
}

// PrepareCallHierarchy implements protocol.Server.
func (*DefaultServer) PrepareCallHierarchy(ctx context.Context, params *protocol.CallHierarchyPrepareParams) (result []protocol.CallHierarchyItem, err error) {
	return nil, ErrNotImplemented
}

// PrepareRename implements protocol.Server.
func (*DefaultServer) PrepareRename(ctx context.Context, params *protocol.PrepareRenameParams) (result *protocol.Range, err error) {
	return nil, ErrNotImplemented
}

// RangeFormatting implements protocol.Server.
func (*DefaultServer) RangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) (result []protocol.TextEdit, err error) {
	return nil, ErrNotImplemented
}

// References implements protocol.Server.
func (*DefaultServer) References(ctx context.Context, params *protocol.ReferenceParams) (result []protocol.Location, err error) {
	return nil, ErrNotImplemented
}

// Rename implements protocol.Server.
func (*DefaultServer) Rename(ctx context.Context, params *protocol.RenameParams) (result *protocol.WorkspaceEdit, err error) {
	return nil, ErrNotImplemented
}

// Request implements protocol.Server.
func (*DefaultServer) Request(ctx context.Context, method string, params interface{}) (result interface{}, err error) {
	return nil, ErrNotImplemented
}

// SemanticTokensFull implements protocol.Server.
func (*DefaultServer) SemanticTokensFull(ctx context.Context, params *protocol.SemanticTokensParams) (result *protocol.SemanticTokens, err error) {
	return nil, ErrNotImplemented
}

// SemanticTokensFullDelta implements protocol.Server.
func (*DefaultServer) SemanticTokensFullDelta(ctx context.Context, params *protocol.SemanticTokensDeltaParams) (result interface{}, err error) {
	return nil, ErrNotImplemented
}

// SemanticTokensRange implements protocol.Server.
func (*DefaultServer) SemanticTokensRange(ctx context.Context, params *protocol.SemanticTokensRangeParams) (result *protocol.SemanticTokens, err error) {
	return nil, ErrNotImplemented
}

// SemanticTokensRefresh implements protocol.Server.
func (*DefaultServer) SemanticTokensRefresh(ctx context.Context) (err error) {
	return ErrNotImplemented
}

// SetTrace implements protocol.Server.
func (*DefaultServer) SetTrace(ctx context.Context, params *protocol.SetTraceParams) (err error) {
	return ErrNotImplemented
}

// ShowDocument implements protocol.Server.
func (*DefaultServer) ShowDocument(ctx context.Context, params *protocol.ShowDocumentParams) (result *protocol.ShowDocumentResult, err error) {
	return nil, ErrNotImplemented
}

// SignatureHelp implements protocol.Server.
func (*DefaultServer) SignatureHelp(ctx context.Context, params *protocol.SignatureHelpParams) (result *protocol.SignatureHelp, err error) {
	return nil, ErrNotImplemented
}

// Symbols implements protocol.Server.
func (*DefaultServer) Symbols(ctx context.Context, params *protocol.WorkspaceSymbolParams) (result []protocol.SymbolInformation, err error) {
	return nil, ErrNotImplemented
}

// TypeDefinition implements protocol.Server.
func (*DefaultServer) TypeDefinition(ctx context.Context, params *protocol.TypeDefinitionParams) (result []protocol.Location, err error) {
	return nil, ErrNotImplemented
}

// WillCreateFiles implements protocol.Server.
func (*DefaultServer) WillCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (result *protocol.WorkspaceEdit, err error) {
	return nil, ErrNotImplemented
}

// WillDeleteFiles implements protocol.Server.
func (*DefaultServer) WillDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (result *protocol.WorkspaceEdit, err error) {
	return nil, ErrNotImplemented
}

// WillRenameFiles implements protocol.Server.
func (*DefaultServer) WillRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (result *protocol.WorkspaceEdit, err error) {
	return nil, ErrNotImplemented
}

// WillSave implements protocol.Server.
func (*DefaultServer) WillSave(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (err error) {
	return ErrNotImplemented
}

// WillSaveWaitUntil implements protocol.Server.
func (*DefaultServer) WillSaveWaitUntil(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (result []protocol.TextEdit, err error) {
	return nil, ErrNotImplemented
}

// WorkDoneProgressCancel implements protocol.Server.
func (*DefaultServer) WorkDoneProgressCancel(ctx context.Context, params *protocol.WorkDoneProgressCancelParams) (err error) {
	return ErrNotImplemented
}

func (s *DefaultServer) Initialize(ctx context.Context, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	return nil, ErrNotImplemented
}

func (s *DefaultServer) Initialized(ctx context.Context, params *protocol.InitializedParams) error {
	return ErrNotImplemented
}

func (s *DefaultServer) Shutdown(ctx context.Context) error {
	return ErrNotImplemented
}

func (s *DefaultServer) Exit(ctx context.Context) error {
	return ErrNotImplemented
}
