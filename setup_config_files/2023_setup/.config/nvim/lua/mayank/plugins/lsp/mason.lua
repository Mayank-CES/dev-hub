-- import mason plugin safely
local mason_status, mason = pcall(require, "mason")
if not mason_status then
  return
end

-- import mason-lspconfig plugin safely
local mason_lspconfig_status, mason_lspconfig = pcall(require, "mason-lspconfig")
if not mason_lspconfig_status then
  return
end

-- import mason-null-ls plugin safely
local mason_null_ls_status, mason_null_ls = pcall(require, "mason-null-ls")
if not mason_null_ls_status then
  return
end

-- enable mason
mason.setup()

mason_lspconfig.setup({
  -- list of servers for mason to install
  ensure_installed = {
    "gopls",        -- Go language server
    "phpactor",     -- PHP language server
    "solargraph",   -- Ruby on Rails language server
    "html",         -- HTML language server
    "cssls",        -- CSS language server
    "lua_ls",       -- Lua language server (for configuring Neovim)
  },
  -- auto-install configured servers (with lspconfig)
  automatic_installation = true, -- not the same as ensure_installed
})

mason_null_ls.setup({
  -- list of formatters & linters for mason to install
  ensure_installed = {
    "gofumpt",      -- Go formatter
    "golangci_lint",-- Go linter
    "phpcs",        -- PHP linter
    "phpcbf",       -- PHP formatter
    -- "rubocop",      -- Ruby formatter and linter
    "prettier",     -- General-purpose formatter (e.g., for HTML, CSS)
    "eslint_d",     -- JavaScript/TypeScript linter
  },
  -- auto-install configured formatters & linters (with null-ls)
  automatic_installation = true,
})
